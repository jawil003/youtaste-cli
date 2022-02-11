import React from "react";

export interface Props {}

/**
 * An Timer React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Timer: React.FC<Props> = ({ children }) => {
  return (
    <div className="absolute top-0 left-0 w-full flex items-center justify-center">
      <div className="rounded-b-lg p-2 shadow-lg font-semibold bg-red-500">
        {children}
      </div>
    </div>
  );
};
