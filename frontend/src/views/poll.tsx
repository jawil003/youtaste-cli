import React from "react";
import { Helmet } from "react-helmet";
import { AnotherPollOption } from "../components/another-poll-option/another-poll-option";
import { Button } from "../components/button/button";
import { PollOption } from "../components/poll-option/poll-option";
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
      <div className="background-card">
        <div className="background-card-title">
          Which Restaurant do you want to order from?
        </div>
        <div className="mt-8 flex flex-col gap-4 ">
          {Object.entries(polls)?.map(([poll, count]) => (
            <PollOption type="button" count={count}>
              {poll}
            </PollOption>
          ))}
          <Button className="mt-8">Add another</Button>
        </div>
      </div>
      <AnotherPollOption />
    </div>
  );
};
