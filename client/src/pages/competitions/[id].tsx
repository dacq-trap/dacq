import { GetServerSideProps } from 'next'
import CompetitionPageBase from '@/components/competitionPageBase'

// 仮の構造
type Props = {
  competitionId: string
}

const Competition = (props: Props) => {
  return (
    <CompetitionPageBase
      competitionId={props.competitionId}
      competitionPageOption='overview'
    >
      問題概要ページ
    </CompetitionPageBase>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const competitionId = ctx.params?.id as string

  const props: Props = {
    competitionId,
  }

  return {
    props: props,
  }
}

export default Competition
