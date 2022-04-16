import axios from "axios";

export const API_URL = "http://localhost:8000/api/v1/";

const $api = axios.create({
  withCredentials: true,
  baseURL: API_URL,
});

$api.interceptors.request.use((config) => {});

export default $api;
