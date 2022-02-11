import React from "react";

export interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  count: number;
}

/**
 * An PollOption React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const PollOption: React.FC<Props> = ({
  children,
  className,
  count,
  ...props
}) => {
  return (
    <button
      type="submit"
      className={
        `relative overflow-hidden flex py-2 px-4  ${`border-white-600 hover:border-white-700 focus:ring-white-500 focus:ring-offset-white-200`} text-black w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg ` +
        className
      }
      {...props}
    >
      {children}
      <div className="absolute top-0 right-0 bottom-0 aspect-square bg-blue-100 rounded-lg h-full flex items-center justify-center">
        {count}
      </div>
    </button>
  );
};
