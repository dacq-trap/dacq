import { GetServerSideProps } from 'next'
import ProblemPageBase from '@/components/ProblemPageBase'

// 仮の構造
type Props = {
  problemId: string
}

const Problem = (props: Props) => {
  return (
    <ProblemPageBase problemId={props.problemId} problemPageOption='overview'>
      問題概要ページ
    </ProblemPageBase>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const problemId = ctx.params?.id as string

  const props: Props = {
    problemId,
  }

  return {
    props: props,
  }
}

export default Problem
