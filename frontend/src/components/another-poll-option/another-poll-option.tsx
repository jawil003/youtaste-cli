import React from "react";
import ReactDOM from "react-dom";
import { FormProvider, useForm } from "react-hook-form";
import { Button } from "../button/button";
import { Input } from "../input/input";
import { Select } from "../select/select";

export interface Props {}

/**
 * An AnotherPollOption React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const AnotherPollOption: React.FC<Props> = () => {
  const methods = useForm({ defaultValues: { provider: "youtaste" } });

  return ReactDOM.createPortal(
    <FormProvider {...methods}>
      <div className="absolute z-50 top-0 left-0 bg-black-20 w-screen h-screen flex items-center justify-center">
        <div className="background-card">
          <div className="background-card-title">Add another Restaurant</div>
          <div className="mt-8">
            <form>
              <Select
                options={[
                  { value: "youtaste", label: "YouTaste" },
                  { value: "lieferando", label: "Lieferando" },
                ]}
                placeholder="YouTaste"
                required
                label="Provider"
                name="provider"
              />
              <div className="mt-2" />
              <Input
                required
                label="Restaurantname"
                placeholder="Restaurant am Markt"
                name="name"
              />

              <Button className=" mt-8">Add</Button>
            </form>
          </div>
        </div>
      </div>
    </FormProvider>,
    document.body
  );
};
