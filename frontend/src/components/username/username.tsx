import React from "react";

export interface Props {}

/**
 * An Username React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Username: React.FC<Props> = ({ children }) => {
  return (
    <span className="flex items-center text-gray-500 text-md">{children}</span>
  );
};
