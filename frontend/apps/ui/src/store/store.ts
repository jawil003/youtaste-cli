import create from "zustand";

interface Store {
  user: { firstname: string; lastname: string };
  setUser: (user: Store["user"]) => void;
}
export const useStore = create<Store>((set) => ({
  user: { firstname: "", lastname: "" },
  setUser: (user: Store["user"]) => set((state) => ({ ...state, user })),
}));
