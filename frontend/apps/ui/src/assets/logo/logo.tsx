import React from "react";

export interface Props {
  className?: string;
}

/**
 * An Logo React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Logo: React.FC<Props> = ({ className }) => {
  return (
    <h1 className={"text-2xl font-bold text-black " + className}>TastyFood</h1>
  );
};
