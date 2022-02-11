import React, { useEffect } from "react";
import { Helmet } from "react-helmet";
import { Button } from "../components/button/button";
import { usePolls } from "../hooks/polls.hook";

export interface Props {}

/**
 * An Poll React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Poll: React.FC<Props> = () => {
  const polls = usePolls();

  return (
    <div className="flex w-full h-full justify-center items-center">
      <Helmet>
        <title>Choose Restaurant | TastyFood</title>
      </Helmet>
      <div className="max-w-md w-full px-4 py-8 bg-white rounded-lg shadow sm:px-6 md:px-8 lg:px-10">
        <div className="self-center text-xl font-light text-gray-600 sm:text-2xl">
          Which Restaurant do you want to order from?
        </div>
        <div className="mt-8 flex flex-col gap-4 ">
          {Object.keys(polls)?.map((poll) => (
            <Button variant="secondary" type="button">
              {poll}
            </Button>
          ))}
          <Button className="mt-8">Add another</Button>
        </div>
      </div>
    </div>
  );
};
