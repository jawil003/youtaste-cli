import React from "react";
import { Link } from "react-router-dom";
import { ReactComponent as CookieIcon } from "../assets/Fortune cookie-cuate.svg";
import { Button } from "../components/button/button";
export interface Props {}

/**
 * An ErrorView React Component.
 * @author
 * @version 0.1
 */
export const ErrorView: React.FC<Props> = () => {
  return (
    <div className="bg-white w-full h-full px-6 mt-6 flex flex-row justify-between items-center relative">
      <div className="flex-1">
        <div className="flex flex-1 m-auto max-w-lg  gap-y-8 justify-center items-start flex-col mb-16">
          <h1 className="font-thin max-w-lg text-6xl text-gray-800">
            Something went wrong
          </h1>
          <Link to={"/"} className="ml-0">
            <Button>Go Back Home</Button>
          </Link>
        </div>
      </div>
      <div className="block flex-1 mx-auto mt-6 md:mt-0 relative">
        <CookieIcon />
      </div>
    </div>
  );
};
