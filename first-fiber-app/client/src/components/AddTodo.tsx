import { useState } from "react";
import { ENDPOINT, TODO } from "../App";
import { KeyedMutator } from "swr";

const AddTodo = ({ mutate }: { mutate: KeyedMutator<TODO[]> }) => {
    const [title, setTitle] = useState('');
    const [body, setBody] = useState('');

    const addHandler = async (values: { title: string, body: string }) => {
        if (values.body.length <= 0 || values.title.length <= 0) {
            return;
        }

        const update = await fetch(`${ENDPOINT}/api/todo`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(values),
        }).then(response => response.json());

        setTitle('');
        setBody('');
        mutate(update);
    }

    return (
        <div>
            <form>
                <input
                    type="text"
                    placeholder="title"
                    value={title}
                    onChange={e => setTitle(e.target.value)}
                />
                <input
                    type="text"
                    placeholder="body"
                    value={body}
                    onChange={e => setBody(e.target.value)}
                />
            </form>
            <button type="submit" onClick={() => addHandler({ title, body })}>Add todo</button>
        </div>
    );
}

export default AddTodo;