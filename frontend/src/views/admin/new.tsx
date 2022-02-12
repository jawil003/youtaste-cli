import React from "react";
import { Helmet } from "react-helmet";
import { FormProvider, useForm } from "react-hook-form";
import { Button } from "../../components/button/button";
import { Input } from "../../components/input/input";
import { ReactComponent as LieferandoLogo } from "../../assets/lieferandoat-small.svg";
import youtasteLogoUrl from "../../assets/youtaste-white-logo.png";

export interface Props {}

/**
 * An AdminNewView React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const AdminNewView: React.FC<Props> = () => {
  const methods = useForm({
    defaultValues: {
      pollTimeout: 0,
      orderTimeout: 0,
    },
  });

  return (
    <FormProvider {...methods}>
      <div className="w-full h-full flex items-center justify-center">
        <Helmet>
          <title>Setup | TastyFood</title>
        </Helmet>
        <div className="background-card">
          <div className="background-card-title">Setup</div>
          <form className="flex flex-col gap-y-2 mt-8">
            <div className="flex gap-x-2 mb-2">
              <Button
                className="bg-orange-400 focus:ring-orange-400"
                type="button"
              >
                <div className="flex gap-x-2 items-center">
                  <LieferandoLogo width={20} />
                  <div className="flex-1">
                    <span>Login</span>
                  </div>
                </div>
              </Button>
              <Button className="bg-red-400 focus:ring-red-400" type="button">
                <div className="flex gap-x-2 items-center">
                  <img src={youtasteLogoUrl} alt="youtasteLogo" width={12} />
                  <div className="flex-1">
                    <span>Login</span>
                  </div>
                </div>
              </Button>
            </div>
            <Input type={"number"} label="Poll Timeout" name="pollTimeout" />
            <Input type={"number"} label="Order Timeout" name="orderTimeout" />
            <Button className="mt-4 ">Submit</Button>
          </form>
        </div>
      </div>
    </FormProvider>
  );
};
