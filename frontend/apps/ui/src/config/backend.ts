import axios from "axios";

export const backend = axios.create({
  baseURL:
    process.env.NODE_ENV === "production"
      ? undefined
      : process.env.REACT_APP_BASE_URL,
  withCredentials: true,
});
