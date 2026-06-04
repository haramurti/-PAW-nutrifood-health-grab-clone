import Navbar from "../components/Navbar";
import CardGrade from "../components/CardGrade";
import Nutrition from "../components/Nutrition";
import FormIngredient from "../components/FormIngredient";
import { useLocation, useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { createProduct, generateNutrition, updateProduct } from "../redux/slices/productSlice";
import { useEffect, useState } from "react";

function CreateProductNutritionPage() {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const { state } = useLocation();
  const { ingredient, message } = useSelector((state) => state.productSlice);
  const { role, token } = useSelector((state) => state.loginSlice);

  const [data, setData] = useState({});
  const [ingredients, setIngredients] = useState([]);
  const [isRegenerating, setIsRegenerating] = useState(false);

  useEffect(() => {
    if (!role || !token) {
      navigate("/");
    }
  }, [role, token, navigate]);

  useEffect(() => {
    if (role && token) {
      setData(ingredient);
      setIngredients(ingredient.ingredients);
      
      if (state && Object.hasOwn(state, "ingredients")) {
        setData(state);
        setIngredients(state.ingredients);
      }
    }
  }, [ingredient, role, token, state]);

  useEffect(() => {
    if (message === 'success generate nutrition') {
      setData(ingredient);
    };
  }, [ingredient, message]);

  const submitHandler = (e) => {
    e.preventDefault();
    
    const payload = {
      ...state,
      ...ingredient,
    };
    
    console.log("HEHEH", state, Object.hasOwn(state, "ingredients"))
    if (Object.hasOwn(state, "ingredients") && message === 'success generate nutrition') {
      dispatch(updateProduct(payload));
    } else if (!Object.hasOwn(state, "ingredients")) {
      dispatch(createProduct(payload));
    } else {
      dispatch(updateProduct(state));
    }

    navigate("/products");
  };

  const regenerateNutrition = async (e) => {
    e.preventDefault();
    setIsRegenerating(true);
    await dispatch(generateNutrition(data.type, ingredients));
    setIsRegenerating(false);
  };

  return (
    <div className="max-w-md min-h-screen mx-auto bg-white shadow-xl relative pb-20">
      {Object.keys(data).length !== 0 ? (
        <div className="relative h-full">
          <Navbar header={"Create Product - Nutri-Score"}></Navbar>
          {console.log("MASUK", data)}
          <div className="mb-6 max-h-[80%] overflow-y-auto">
            <div className="input-group-uniq">
              <label htmlFor="name" className="product-label-uniq">
                Product Name
              </label>
              <input
                type="text"
                id="name"
                className="product-input-uniq"
                value={state.name}
                readOnly
                disabled
              />
            </div>

            <div className="mx-4 mb-3">
              <CardGrade nutrition={data} />
            </div>

            <div className="mx-4 mb-6">
              <Nutrition nutrition={data.nutrition} />
            </div>

            <div className="mx-4 mb-3">
              <FormIngredient
                ingredient={ingredients}
                setIngredients={setIngredients}
              />
            </div>
          </div>

          <div className="absolute bottom-0 w-full text-sm bg-white">
            <div className="flex gap-3 pt-3 pb-5 mx-4">
              <button
                className={`bg-white grow outline outline-1 outline-[var(--secondary)] rounded py-2 ${isRegenerating ? 'opacity-50 cursor-not-allowed' : ''}`}
                onClick={regenerateNutrition}
                disabled={isRegenerating}
              >
                {isRegenerating ? "Sedang diproses... ⏳" : "Generate Nutri-Score"}
              </button>
              <button
                className="bg-[var(--primary)] w-1/3 rounded py-2 text-white"
                onClick={submitHandler}
              >
                Submit
              </button>
            </div>
          </div>
        </div>
      ) : message === 'error' ? (
        <div className="flex flex-col items-center justify-center h-full pt-20">
          <h1 className="text-red-500 font-bold text-xl mb-4">Gagal Memproses Data</h1>
          <p className="text-gray-600 mb-6">Terjadi kesalahan saat memproses data. Silakan coba lagi.</p>
          <button 
            onClick={() => navigate(-1)}
            className="bg-[var(--primary)] text-white px-6 py-2 rounded-lg hover:bg-opacity-90"
          >
            Kembali & Coba Lagi
          </button>
        </div>
      ) : (
        <div className="flex flex-col items-center justify-center h-full pt-20">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-[var(--primary)] mb-4"></div>
          <h1 className="text-lg font-medium text-gray-700">Sedang diproses, mohon tunggu sebentar... ⏳</h1>
        </div>
      )}
    </div>
  );
}

export default CreateProductNutritionPage;
