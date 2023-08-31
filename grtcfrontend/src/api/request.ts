import Request from "./http";
const baseURL = 'http://localhost:8000/api/v1'
//import.meta.env.VITE_APP_NODE_ENV === 'production' ? import.meta.env.VITE_APP_BASE_API :
const request = new Request({
  baseURL,
  timeout: 20000
})

export default request