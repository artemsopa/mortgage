import axios from "axios";

axios.defaults.baseURL = "http://localhost:8000/api/v1/"

axios.interceptors.response.use(resp => resp, async error => {
    if (error.response.status === 401) {
  
    }
    return error;
});
