import React, { useState } from "react";
import { Helmet } from "react-helmet";
import { AnotherPollOption } from "../components/another-poll-option/another-poll-option";
import { Button } from "../components/button/button";
import { PollOption } from "../components/poll-option/poll-option";
import { ProviderSidebar } from "../components/provider-sidebar/provider-sidebar";
import { Timer } from "../components/timer/timer";
import { usePolls } from "../hooks/polls.hook";
import PollService from "../services/poll.service";
import { useTranslation } from "react-i18next";
export interface Props {}

/**
 * An Poll React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Poll: React.FC<Props> = () => {
  const polls = usePolls();
  const { t } = useTranslation("poll");

  const [open, setOpen] = useState(false);

  return (
    <div className="flex w-full h-full justify-center items-center">
      <Timer />
      <ProviderSidebar />
      <Helmet>
        <title>{t("headline")} | TastyFood</title>
      </Helmet>
      <div className="flex flex-col background-card max-h-3/4">
        <div className="background-card-title mb-8">{t("headline")}</div>
        <div className="flex-1 overflow-auto no-scrollbar p-2">
          <div className="flex flex-col gap-4 flex-1">
            {Object.entries(polls)?.map(([poll, count]) => (
              <PollOption
                onClick={async () => {
                  const pollService = new PollService();
                  await pollService.create({ restaurantName: poll });
                }}
                type="button"
                count={count}
              >
                {poll}
              </PollOption>
            ))}
          </div>
        </div>
        <Button onClick={() => setOpen(true)} className="mt-8">
          {t("addAnother")}
        </Button>
      </div>
      <AnotherPollOption open={open} onClose={() => setOpen(false)} />
    </div>
  );
};
