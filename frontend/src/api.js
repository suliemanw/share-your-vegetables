import axios from 'axios'

const host = window.location.hostname;

const api = axios.create({
  baseURL: `http://${host}:8081`,
  timeout: 30000
});

export default api;
