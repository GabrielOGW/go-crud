import { useState } from "react";
import { useForm } from "@mantine/form";
import { Button, Group, Modal, Textarea, TextInput } from "@mantine/core";
import { ENDPOINT, Todo } from "../App";
import { KeyedMutator } from "swr";

export default function AddTodo({ mutate }: { mutate: KeyedMutator<Todo[]> }) {
  const [isOpen, setIsOpen] = useState(false);

  const form = useForm({
    initialValues: {
      title: "",
      body: "",
    },
  });

  async function createTodo(values: { title: string; body: string }) {
    const update = await fetch(`${ENDPOINT}/api/todos`, {
      method: "POST",
      headers: {
        "Content-type": "application/json",
      },
      body: JSON.stringify(values),
    }).then((r) => r.json());

    mutate(update);
    form.reset();
    setIsOpen(false);
  }

  return (
    <>
      <Modal
        opened={isOpen}
        onClose={() => setIsOpen(false)}
        title="Create Todo"
      >
        <form onSubmit={form.onSubmit(createTodo)}>
          <TextInput
            required
            mb={12}
            label="Afazer"
            placeholder="Oque você precisa fazer?"
            {...form.getInputProps("title")}
          />
          <Textarea
            required
            mb={12}
            label="Descrição"
            placeholder="Quais são os detalhes da tarefa?"
            {...form.getInputProps("body")}
          />
          <Button type="submit">CRIAR AFAZER</Button>
        </form>
      </Modal>

      <Group position="center">
        <Button fullWidth mb={12} onClick={() => setIsOpen(true)}>
          ADICIONAR AFAZER
        </Button>
      </Group>
    </>
  );
}
