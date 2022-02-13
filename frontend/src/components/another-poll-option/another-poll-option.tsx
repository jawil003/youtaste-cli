import React from "react";
import ReactDOM from "react-dom";
import { FormProvider, useForm } from "react-hook-form";
import { Button } from "../button/button";
import { Input } from "../input/input";
import { Select } from "../select/select";
import { XIcon } from "@heroicons/react/solid";
import { ProviderSidebar } from "../provider-sidebar/provider-sidebar";
import PollService from "../../services/poll.service";
export interface Props {
  open: boolean;
  onClose: () => void;
}

/**
 * An AnotherPollOption React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const AnotherPollOption: React.FC<Props> = ({ onClose, open }) => {
  const methods = useForm({ defaultValues: { provider: "youtaste" } });

  const handleSubmit = async (values: { name?: string; provider: string }) => {
    const pollService = new PollService();

    await pollService.create({
      restaurantName: values.name ?? "",
      provider: values.provider as "youtaste" | "lieferando",
    });

    methods.reset();
    onClose();
  };

  if (open)
    return ReactDOM.createPortal(
      <FormProvider {...methods}>
        <div className="absolute  z-10 top-0 left-0 bg-black-60 w-screen h-screen flex items-center justify-center">
          <ProviderSidebar />
          <div className="background-card relative">
            <button
              onClick={() => onClose()}
              className="absolute top-3 right-3"
            >
              <XIcon width="1.25rem" />
            </button>
            <div className="background-card-title">Add another Restaurant</div>
            <div className="mt-8">
              <form onSubmit={methods.handleSubmit(handleSubmit)} noValidate>
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
      document.getElementById("modal") as HTMLElement
    );
  return null;
};
