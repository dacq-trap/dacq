import { GetServerSideProps } from 'next'
import CompetitionPageBase from '@/components/competitionPageBase'

// 仮の構造
type Props = {
  competitionId: string
  competitionName: string
}

const Leaderboard = (props: Props) => {
  return (
    <CompetitionPageBase
      // TODO: コンテキスト情報の整理
      competitionId={props.competitionId}
      competitionName={props.competitionName}
      isCompetitionRegistered={true}
      competitionPageOption='leaderboard'
    >
      順位表ページ
    </CompetitionPageBase>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const competitionId = ctx.params?.id as string
  const competitionName = ''
  const props: Props = {
    competitionId,
    competitionName,
  }

  return {
    props: props,
  }
}

export default Leaderboard
