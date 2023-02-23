// "use client"
import { GetServerSideProps } from "next"
import Head from "next/head"
import { useEffect, useState } from "react"
import { GripVertical } from "tabler-icons-react"

type HelloResponse = { message: string }

const fetchFromClient = async (): Promise<HelloResponse> => {
  const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/hello`)
  const data = await res.json()
  return data as HelloResponse
}

function Home({ data }: { data: HelloResponse }) {
  const [hello, setHello] = useState<string>("")

  useEffect(() => {
    fetchFromClient()
      .then((res) => setHello(res.message))
      .catch(() => setHello("Failed to fetch from client :("))
  }, [])

  return (
    <>
      <Head>
        <title>Quizaction</title>
        <meta name="description" content="Quiz app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <div className="flex min-h-screen items-center justify-center">
          <div className="group flex flex-row items-center gap-2 p-24">
            <GripVertical
              size={58}
              strokeWidth={3}
              className="text-black drop-shadow-md transition-colors duration-700 group-hover:text-green-300"
            />
            <div>
              <h1 className="flex text-3xl font-bold drop-shadow-md">
                <span className="transition-colors duration-700 group-hover:text-green-300">
                  Quiz
                </span>
                action
              </h1>
              <p>{data.message}</p>
            </div>
            <small className="rotate-45 opacity-25">{hello}</small>
          </div>
        </div>
      </main>
    </>
  )
}

export const getServerSideProps: GetServerSideProps = async () => {
  const backendUrl = process.env.BACKEND_URL
  const res = await fetch(`${backendUrl}/hello`)
  const data = (await res.json()) as HelloResponse
  return {
    props: {
      data,
    },
  }
}

export default Home
