import CardComponent from "@/components/molecules/CardComponent";
import InputForm from "@/components/molecules/InputForm";
import { Button } from "@/components/ui/button";
import React from "react";

const ExpenseCard = () => {
  return (
    <CardComponent title="Add Expense">
      <form action="" className="flex flex-col gap-4">
        <InputForm
          id="amount"
          name="amount"
          type="number"
          placeholder="ex : 10000, ">
          amount
        </InputForm>
        <InputForm
          id="notes"
          name="notes"
          type="text"
          placeholder="ex : buy a new car, buy a new house ">
          notes
        </InputForm>
        <Button type="submit" className="bg-teal-700">
          Submit
        </Button>
      </form>
    </CardComponent>
  );
};

export default ExpenseCard;
