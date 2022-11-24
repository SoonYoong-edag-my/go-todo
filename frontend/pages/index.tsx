import type {GetServerSideProps, NextPage} from 'next'
import Head from 'next/head'
import { useRouter } from 'next/router';
import { useState } from 'react';

interface Todo {
  id: number;
  name: string;
  completed: boolean;
}

interface HomeProps {
  data: Todo[]
}

const Home: NextPage<HomeProps> = ({ data }) => {
  const router = useRouter();
  const [name, setName] = useState("");

  const createTodo = () => {
    fetch(`http://localhost:8080/todos`,{
      body: JSON.stringify({ name, completed: false }),
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST'
    }).then(() => {
      router.replace(router.asPath);
      setName("");
    })
  }

  const updateTodo = (todo: Todo, completed: boolean) => {
    fetch(`http://localhost:8080/todos/${todo.id}`,{
      body: JSON.stringify({ ...todo, completed }),
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'PUT'
    }).then(() => {
      router.replace(router.asPath);
    })
  }

  const deleteTodo = (id: number) => {
    fetch(`http://localhost:8080/todos/${id}`,{
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'DELETE'
    }).then(() => {
      router.replace(router.asPath);
    })
  }

  console.log(data);
  return (
    <div className="flex min-h-screen flex-col py-2">
      <Head>
        <title>ToDo App</title>
        <link rel="icon" href="/favicon.ico"/>
      </Head>
      <div className="h-screen w-full flex justify-center bg-gray-100 font-sans">
        <div className="bg-white rounded shadow p-6 m-4 w-full lg:w-3/4 lg:max-w-lg">
          <div className="mb-4">
            <h1 className="text-grey-darkest underline">Todo List</h1>
            <div className="flex mt-4">
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-grey-darker"
                value={name}
                onChange={(event: any) => setName(event.target.value)}
                placeholder="Add Todo"/>
              <button
                className="flex-no-shrink p-2 border-2 rounded text-teal-400 border-teal-400 hover:text-white hover:bg-teal-400"
                onClick={createTodo}>Add
              </button>
            </div>
          </div>
          <div>
            {data.map(todo => (
              <div className="flex mb-4 items-center">
                {todo.completed ? 
                <>
                  <p className="w-full line-through text-green">{todo.name}</p>
                  <button
                    className="flex-no-shrink p-2 ml-4 mr-2 border-2 rounded hover:text-white text-gray-400 border-grey-400 hover:bg-gray-400"
                    onClick={() => updateTodo(todo, false)}>Undone
                  </button>
                </>
                :
                <>
                  <p className="w-full text-grey-darkest">{todo.name}</p>
                  <button
                    className="flex-no-shrink p-2 ml-4 mr-2 border-2 rounded hover:text-white text-green-400 border-green-400 hover:bg-green-400"
                    onClick={() => updateTodo(todo, true)}>Done
                  </button>
                </>
                }
              <button
                className="flex-no-shrink p-2 ml-2 border-2 rounded text-red-400 border-red-400 hover:text-white hover:bg-red-400"
                onClick={() => deleteTodo(todo.id)}>Remove
              </button>
            </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}


// This gets called on every request
export const getServerSideProps: GetServerSideProps  =  async (context) => {
  // Fetch data from external API
  const res = await fetch(`http://localhost:8080/todos`)
  const data = await res.json()

  // Pass data to the page via props
  return { props: { data } }
}

export default Home
