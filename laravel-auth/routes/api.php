<?php

use App\Http\Controllers\AuthController;
use Illuminate\Support\Facades\Route;

// Public routes
Route::prefix('auth')->group(function () {
    Route::post('/register',      [AuthController::class, 'register']);
    Route::post('/login',         [AuthController::class, 'login']);
    Route::get('/session-check',  [AuthController::class, 'sessionCheck']);
});

// Protected routes
Route::prefix('auth')->middleware('auth.jwt')->group(function () {
    Route::post('/logout', [AuthController::class, 'logout']);
    Route::get('/me',      [AuthController::class, 'me']);
});
