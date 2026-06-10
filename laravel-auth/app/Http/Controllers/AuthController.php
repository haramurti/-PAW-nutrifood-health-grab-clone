<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;
use Tymon\JWTAuth\Facades\JWTAuth;
use Tymon\JWTAuth\Exceptions\JWTException;
use Illuminate\Support\Facades\Cookie;

class AuthController extends Controller
{
    // ── Register ──────────────────────────────────────────────────────────────

    public function register(Request $request): JsonResponse
    {
        $validator = Validator::make($request->all(), [
            'name'     => 'required|string|max:255',
            'email'    => 'required|string|email|max:255|unique:users',
            'password' => 'required|string|min:6|confirmed',
        ]);

        if ($validator->fails()) {
            return response()->json([
                'message' => 'Validation error',
                'errors'  => $validator->errors(),
            ], 422);
        }

        $user = User::create([
            'name'     => $request->name,
            'email'    => $request->email,
            'password' => Hash::make($request->password),
        ]);

        $token = JWTAuth::fromUser($user);

        return response()->json([
            'message' => 'success',
            'data'    => [
                'user'  => $user,
                'token' => $token,
            ],
        ], 201);
    }

    // ── Login ─────────────────────────────────────────────────────────────────

    public function login(Request $request): JsonResponse
    {
        $credentials = $request->only('email', 'password');

        try {
            if (!$token = JWTAuth::attempt($credentials)) {
                return response()->json([
                    'message' => 'Invalid email or password',
                ], 401);
            }
        } catch (JWTException $e) {
            return response()->json([
                'message' => 'Could not create token',
                'error'   => $e->getMessage(),
            ], 500);
        }

        $user = JWTAuth::user();

        // Set session
        session([
            'loggedin'  => true,
            'username'  => $user->name,
            'user_id'   => $user->id,
        ]);

        // Set cookie role (berlaku 1 hari)
        $cookie = Cookie::make('user_role', $request->input('role', 'customer'), 60 * 24);

        return response()->json([
            'message' => 'success',
            'data'    => [
                'user'  => $user,
                'token' => $token,
            ],
        ])->withCookie($cookie);
    }

    // ── Logout ────────────────────────────────────────────────────────────────

    public function logout(): JsonResponse
    {
        try {
            JWTAuth::invalidate(JWTAuth::getToken());
        } catch (JWTException $e) {
            return response()->json([
                'message' => 'Failed to logout',
                'error'   => $e->getMessage(),
            ], 500);
        }

        // Clear session dan hapus cookie
        session()->flush();

        return response()->json([
            'message' => 'success',
        ])->withCookie(Cookie::forget('user_role'));
    }

    // ── Session & Cookie check ────────────────────────────────────────────────

    public function sessionCheck(Request $request): JsonResponse
    {
        return response()->json([
            'message' => 'success',
            'data'    => [
                'session' => [
                    'loggedin' => session('loggedin', false),
                    'username' => session('username', null),
                    'user_id'  => session('user_id', null),
                ],
                'cookie' => [
                    'user_role' => $request->cookie('user_role', null),
                ],
            ],
        ]);
    }

    // ── Me (get current user) ─────────────────────────────────────────────────

    public function me(): JsonResponse
    {
        try {
            $user = JWTAuth::parseToken()->authenticate();
        } catch (JWTException $e) {
            return response()->json(['message' => 'Unauthorized'], 401);
        }

        return response()->json([
            'message' => 'success',
            'data'    => $user,
        ]);
    }
}
