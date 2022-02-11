import React, { useEffect } from "react";
import { Helmet } from "react-helmet";
import { FormProvider, useForm, useWatch } from "react-hook-form";
import { useQueryClient } from "react-query";
import { useNavigate, useParams } from "react-router-dom";
import { Badge } from "../components/badge/badge";
import { Button } from "../components/button/button";
import { Input } from "../components/input/input";
import { Routes } from "../enums/routes.enum";
import { useUser } from "../hooks/user.hook";
import OrderService from "../services/order.service";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";

export interface Props {}

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
export const CreateOrderView: React.FC<Props> = () => {
  const { name } = useParams<"name">();

  const methods = useForm<FormData>({
    defaultValues: { mealName: "", variant: "", variants: [] },
    resolver: yupResolver(schema),
  });

  useEffect(() => {
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

  const { data: user } = useUser();

  const navigate = useNavigate();

  const variant = useWatch({ name: "variant", control: methods.control });
  const variants = useWatch({ name: "variants", control: methods.control });

  const queryClient = useQueryClient();

  const onSubmit = async (value: FormData) => {
    const orderService = new OrderService();

    await orderService.createOrUpdate([{ name: value.mealName, variants }]);

    await queryClient.invalidateQueries(["orders-by-user"]);

    navigate(
      Routes.ORDER_CONFIRM.replace(
        ":user",
        `${user?.firstname.toLowerCase()}_${user?.lastname.toLowerCase()}`
      )
    );
    methods.reset();
  };

  return (
    <div className="w-screen h-screen flex items-center justify-center">
      <Helmet>
        <title>New Order | TastyFood</title>
      </Helmet>
      <div className="flex flex-col background-card">
        <div className="background-card-title">What do you like to eat?</div>
        <div className="mt-8">
          <FormProvider {...methods}>
            <form noValidate onSubmit={methods.handleSubmit(onSubmit)}>
              <Input
                required
                className="mb-2"
                name="mealName"
                placeholder="Pizza Cipola"
                label="Mealname"
              />
              <div className="flex">
                <Input
                  className="flex-none"
                  name="variant"
                  placeholder="Big"
                  label="Variant"
                />
                <div className="p-6">
                  <Button
                    onClick={() => {
                      if (!variant) {
                        return;
                      }
                      methods.setValue("variants", [...variants, variant]);
                      methods.setValue("variant", "");
                    }}
                    className="flex-auto"
                    type="button"
                  >
                    Add
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
              <Button className="mb-2">Submit</Button>
            </form>
          </FormProvider>
        </div>
      </div>
    </div>
  );
};
