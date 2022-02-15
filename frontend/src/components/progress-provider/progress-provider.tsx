import React, { useEffect } from "react";
import { useProgress } from "../../hooks/useProgress.hook";
import { useUser } from "../../hooks/user.hook";

export interface Props {}

export const ProgressProvider = React.createContext<string | null>(null);

/**
 * An ProgressProvider React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const ProgressProviderWrapper: React.FC<Props> = ({ children }) => {
  const { progress, refetch } = useProgress();
  const { data: user, isFetched } = useUser();

  useEffect(() => {
    if (user && isFetched) refetch();
  }, [refetch, user, isFetched]);

  return (
    <ProgressProvider.Provider value={progress}>
      {children}
    </ProgressProvider.Provider>
  );
};
