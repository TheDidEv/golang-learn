import useSWR from 'swr';
import './App.css';
import AddTodo from './components/AddTodo';

export const ENDPOINT = "http://127.0.0.1:4000";


export interface TODO {
  id: number;
  title: string;
  body: string;
  done: boolean;
}

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((response) => response.json());

function App() {
  const { data, mutate } = useSWR<TODO[]>('api/todo', fetcher);

  const markTodo = async (id: number) => {
    const update = await fetch(`${ENDPOINT}/api/todo/${id}/done`, {
      method: "PATCH",
    }).then(responce => responce.json());

    mutate(update);
  }

  return (
    <div>
      <div>
        {data?.map(todo => (
          <div key={`todo__list__${todo.id}`}>
            <input
              type="checkbox"
              checked={todo.done}
              onChange={() => markTodo(todo.id)}
            />
            <div>{todo.title}</div>
            <div>{todo.body}</div>
          </div>
        ))}
      </div>
      <AddTodo mutate={mutate} />
    </div>
  );
}

export default App
