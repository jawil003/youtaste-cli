import React, { useState } from "react";
import { Helmet } from "react-helmet";
import { FormProvider, useForm } from "react-hook-form";
import { useQueryClient } from "react-query";
import { useNavigate } from "react-router-dom";
import { Button } from "../components/button/button";
import { CreateOrderView } from "../components/create-order/create-order";
import { OrderList } from "../components/order-list/order-list";
import { Timer } from "../components/timer/timer";
import {} from "../enums/routes.enum";
import { useOrdersByUser } from "../hooks/ordersByUser.hook";
import OrderService from "../services/order.service";

export interface Props {}

/**
 * An OrderConfirmation React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const OrderConfirmation: React.FC<Props> = () => {
  const { data: result } = useOrdersByUser();
  const queryClient = useQueryClient();
  const methods = useForm();

  const [open, setOpen] = useState(false);
  const [name, setName] = useState<string | undefined>(undefined);

  return (
    <FormProvider {...methods}>
      <Helmet>
        <title>My Orders | TastyFood</title>
      </Helmet>
      <div className="flex items-center justify-center w-full h-full">
        <Timer></Timer>
        <OrderList
          headline="My Orders"
          items={
            result?.orders?.map(({ name, variants }) => ({
              headline: name ?? "",
              description: variants?.join(", ") ?? "",
              onEditClick: () => {
                setName(name);
                setOpen(true);
              },
              onDeleteClick: async () => {
                const orderService = new OrderService();

                if (!name) return;

                await orderService.deleteOrder(name);

                await queryClient.invalidateQueries(["orders-by-user"]);
              },
              size: 1,
            })) ?? []
          }
        >
          <Button
            className="mt-4"
            type="button"
            onClick={() => {
              setName(undefined);
              setOpen(true);
            }}
          >
            Add
          </Button>
        </OrderList>
      </div>
      <CreateOrderView
        open={open}
        name={name}
        onClose={() => {
          setOpen(false);
        }}
      />
    </FormProvider>
  );
};
