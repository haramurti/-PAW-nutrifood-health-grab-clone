import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const productSlice = createSlice({
  name: 'product',
  initialState: {
    products: [],
    product: {},
    url: '',
    ingredient: {},
    message: ''
  },
  reducers: {
    setAllProducts(state, action) {
      state.products = action.payload;
    },
    setDetailProduct(state, action) {
      state.product = action.payload;
    },
    setImageUrl(state, action) {
      state.url = action.payload;
    },
    setIngredient(state, action) {
      state.ingredient = action.payload.data;
      state.message = action.payload.message;
    },
    setMessage(state, action) {
      state.message = action.payload;
    }
  }
});

export const { setAllProducts, setDetailProduct, setImageUrl, setIngredient, setMessage } = productSlice.actions;

export const getAllProducts = (search = '') => async (dispatch, getState) => {
  try {
    const token = getState().loginSlice.token;
    const { data: { data } } = await axios({
      method: 'get',
      url: `http://localhost:8000/api/product?search=${search}`,
      headers: {
        Authorization: `Bearer ${token}`
      },
      responseType: 'json'
    });

    data.forEach(v => {
      v.price = Intl.NumberFormat("id-ID", {
        style: "currency",
        currency: "IDR"
      }).format(v.price).slice(0, -3);
    });

    return dispatch(setAllProducts(data));
  } catch (err) {
    return dispatch(setAllProducts([]));
  };
};

export const getDetailProduct = (id) => async (dispatch, getState) => {
  try {
    const token = getState().loginSlice.token;
    const { data: { data } } = await axios({
      method: 'get',
      url: `http://localhost:8000/api/product/${id}`,
      headers: {
        Authorization: `Bearer ${token}`
      },
      responseType: 'json'
    });

    return dispatch(setDetailProduct(data));
  } catch (err) {
    return dispatch(setDetailProduct({}));
  };
};

export const uploadImage = (image) => async (dispatch, getState) => {
  try {
    const token = getState().loginSlice.token;
    const formData = new FormData();
    formData.append('image', image);
    const { data: { data: { url } } } = await axios({
      method: 'post',
      url: `http://localhost:8000/api/upload`,
      data: formData,
      headers: {
        Authorization: `Bearer ${token}`
      },
      responseType: 'json'
    });

    return dispatch(setImageUrl(`http://localhost:8000${url}`));
  } catch (err) {
    return dispatch(setImageUrl(''));
  };
};

export const generateIngredient = (name) => async (dispatch, getState) => {
  try {
    const token = getState().loginSlice.token;
    const { data: { data } } = await axios({
      method: 'post',
      url: `http://localhost:8000/api/product/generate/ingredient-nutriscore`,
      data: { name },
      headers: {
        Authorization: `Bearer ${token}`
      },
      responseType: 'json'
    });

    console.log(data);
    return dispatch(setIngredient({ data, message: 'success generate ingredient' }));
  } catch (err) {
    return dispatch(setIngredient({ data: {}, message: 'error' }));
  };
};

export const createProduct = (product) => async (dispatch, getState) => {
  try {
    const token = getState().loginSlice.token;
    const { data: { message } } = await axios({
      method: 'post',
      url: `http://localhost:8000/api/product`,
      data: product,
      headers: {
        Authorization: `Bearer ${token}`
      },
      responseType: 'json'
    });

    return dispatch(setMessage(message));
  } catch (err) {
    return dispatch(setMessage('error'));
  };
};

export const generateNutrition = (type, ingredient) => async (dispatch, getState) => {
  try {
    const token = getState().loginSlice.token;
    const { data: { data } } = await axios({
      method: 'post',
      url: `http://localhost:8000/api/product/generate/nutriscore`,
      data: {
        type,
        ingredients: ingredient
      },
      headers: {
        Authorization: `Bearer ${token}`
      },
      responseType: 'json'
    });

    return dispatch(setIngredient({ data, message: 'success generate nutrition' }));
  } catch (err) {
    return dispatch(setIngredient({ data: {}, message: 'error' }));
  };
};

export const updateProduct = (product) => async (dispatch, getState) => {
  try {
    const token = getState().loginSlice.token;
    const { data: { message } } = await axios({
      method: 'put',
      url: `http://localhost:8000/api/product/${product.id}`,
      data: product,
      headers: {
        Authorization: `Bearer ${token}`
      },
      responseType: 'json'
    });

    console.log(message);
    return dispatch(setMessage(message));
  } catch (err) {
    return dispatch(setMessage('error'));
  };
};

export default productSlice.reducer;
