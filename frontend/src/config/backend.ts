import axios from "axios";

export const backend = axios.create({ baseURL: process.env.BASE_URL });
