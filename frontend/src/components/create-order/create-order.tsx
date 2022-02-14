import React, { useEffect } from "react";
import { FormProvider, useForm, useWatch } from "react-hook-form";
import { useQueryClient } from "react-query";
import { Badge } from "../badge/badge";
import { Button } from "../button/button";
import { Input } from "../input/input";
import OrderService from "../../services/order.service";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import ReactDOM from "react-dom";
import { XIcon } from "@heroicons/react/solid";
import { Queries } from "../../enums/queries.enum";
import { useTranslation } from "react-i18next";
import { PlusIcon } from "@heroicons/react/solid";
export interface Props {
  open: boolean;
  name?: string;
  onClose: () => void;
}

interface FormData {
  mealName: string;
  variant: string;
  variants: string[];
}

const schema = yup.object({
  mealName: yup.string().required("Mealname is required"),
});

/**
 * An CreateOrderView React Component.
 * @author
 * @version 0.1
 */
export const CreateOrderView: React.FC<Props> = ({ open, name, onClose }) => {
  const methods = useForm<FormData>({
    defaultValues: { mealName: "", variant: "", variants: [] },
    resolver: yupResolver(schema),
  });

  const { t } = useTranslation("order-confirmation");

  useEffect(() => {
    console.log(name);
    if (name) {
      (async () => {
        const orderService = new OrderService();

        const { data } = await orderService.getByUserAndName(name);

        if (!data?.order) return;

        methods.setValue("mealName", data.order.name ?? "");
        methods.setValue("variants", data.order.variants ?? []);
      })();
    }
  }, [methods, name]);

  const variant = useWatch({ name: "variant", control: methods.control });
  const variants = useWatch({ name: "variants", control: methods.control });

  const queryClient = useQueryClient();

  const onSubmit = async (value: FormData) => {
    const orderService = new OrderService();

    await orderService.createOrUpdate([{ name: value.mealName, variants }]);

    await queryClient.invalidateQueries(Queries.ORDERS_BY_USER);

    methods.reset();
    onClose();
  };

  return ReactDOM.createPortal(
    open && (
      <div className="w-screen h-screen flex items-center justify-center bg-black-60">
        <div className="flex flex-col background-card relative">
          <button onClick={() => onClose()} className="absolute top-3 right-3">
            <XIcon width="1.25rem" />
          </button>
          <div className="background-card-title">{t("addButton")}</div>
          <div className="mt-8">
            <FormProvider {...methods}>
              <form noValidate onSubmit={methods.handleSubmit(onSubmit)}>
                <Input
                  required
                  className="mb-2"
                  name="mealName"
                  placeholder={t("mealNamePlaceholder")}
                  label={t("mealName")}
                />
                <div className="flex">
                  <Input
                    className="flex-none"
                    name="variant"
                    placeholder={t("variantPlaceholder")}
                    label={t("variant")}
                  />
                  <div className="p-6">
                    <Button
                      variant="secondary"
                      onClick={() => {
                        if (!variant) {
                          return;
                        }
                        methods.setValue("variants", [...variants, variant]);
                        methods.setValue("variant", "");
                      }}
                      className="flex-auto h-10"
                      type="button"
                    >
                      <PlusIcon width={"1.25rem"} />
                    </Button>
                  </div>
                </div>
                <div className="flex gap-1 mb-4">
                  {variants?.map((variant, index) => (
                    <Badge
                      onClick={() => {
                        const newVariants = variants.filter(function (_, i) {
                          return i !== index;
                        });

                        methods.setValue("variants", newVariants);
                      }}
                    >
                      {variant}
                    </Badge>
                  ))}
                </div>
                <Button className="mb-2">{t("addButton")}</Button>
              </form>
            </FormProvider>
          </div>
        </div>
      </div>
    ),
    document.getElementById("modal") as HTMLElement
  );
};
