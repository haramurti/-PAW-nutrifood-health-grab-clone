import { faArrowLeft, faRightFromBracket } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { logout } from "../redux/slices/loginSlice";

function Navbar({ header, showBack = true }) {
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const token = useSelector((state) => state.loginSlice.token);

  const handleLogout = () => {
    dispatch(logout());
    navigate("/");
  };

  return (
    <div className="flex items-center justify-between bg-white drop-shadow-md py-3 px-5 text-[var(--neutral)] mb-5">
      <div className="flex items-center gap-3">
        {showBack && (
          <FontAwesomeIcon 
            icon={faArrowLeft} 
            className="cursor-pointer text-sm hover:text-[var(--primary)] transition-all mr-2" 
            onClick={() => navigate(-1)} 
          />
        )}
        <h1 className="font-semibold text-base">{header}</h1>
      </div>
      {token && (
        <button 
          onClick={handleLogout}
          className="flex items-center gap-1.5 px-3 py-1.5 text-xs font-semibold text-red-600 bg-red-50 hover:bg-red-100 rounded-lg cursor-pointer transition-all border border-red-200"
        >
          <FontAwesomeIcon icon={faRightFromBracket} />
          Logout
        </button>
      )}
    </div>
  );
}

export default Navbar;