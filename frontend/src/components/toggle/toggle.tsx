import React from "react";
import { useFormContext, useWatch } from "react-hook-form";

export interface Props {
  name: string;
  className?: string;
}

/**
 * An Toggle React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Toggle: React.FC<Props> = ({ children, name, className }) => {
  const { register } = useFormContext();

  const value = useWatch({ name }) as boolean;

  return (
    <div className={className}>
      <label className="flex relative items-center mb-4 cursor-pointer">
        <input
          {...register(name)}
          type="checkbox"
          className="sr-only"
          checked={value}
        />
        <div className="w-11 h-6 bg-gray-200 rounded-full border border-gray-200 toggle-bg dark:bg-gray-700 dark:border-gray-600"></div>
        <span className="ml-3 text-sm font-medium text-gray-900 dark:text-gray-300">
          {children}
        </span>
      </label>
    </div>
  );
};
