import React from "react";
import { useProgress } from "../../hooks/useProgress.hook";

export interface Props {}

export const ProgressProvider = React.createContext<string | null>(null);

/**
 * An ProgressProvider React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const ProgressProviderWrapper: React.FC<Props> = ({ children }) => {
  const { progress } = useProgress();

  return (
    <ProgressProvider.Provider value={progress}>
      {children}
    </ProgressProvider.Provider>
  );
};
