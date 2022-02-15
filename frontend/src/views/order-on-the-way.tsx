import React from "react";
import { Link } from "react-router-dom";
import { Button } from "../components/button/button";
import { ReactComponent as DeliveryGuy } from "../assets/delivery-animate.svg";
import { Helmet } from "react-helmet";
import { useTranslation } from "react-i18next";
export interface Props {}

/**
 * An OrderOnTheWayView React Component.
 * @author
 * @version 0.1
 */
export const OrderOnTheWayView: React.FC<Props> = () => {
  const { t } = useTranslation("on-the-way");

  return (
    <div className="bg-white w-full h-full px-12 mt-6 flex flex-row justify-between items-center relative">
      <Helmet>
        <title>{t("headline")} | TastyFood</title>
      </Helmet>
      <div className="flex-1">
        <div className="flex flex-1 m-auto max-w-lg  justify-center items-start flex-col mb-16">
          <h1 className="font-thin max-w-lg text-6xl text-gray-800">
            {t("yourOrderIsOnTheWay")}
          </h1>
          <p className="font-light text-xl text-gray-800 mt-4">
            {t("moreInformation")}
          </p>
          <Link to={"/"} className="ml-0 mt-8">
            <Button>{t("moreInfosButton")}</Button>
          </Link>
        </div>
      </div>
      <div className="block flex-1 mx-auto mt-6 md:mt-0 relative">
        <DeliveryGuy height={"100%"} />
      </div>
    </div>
  );
};
