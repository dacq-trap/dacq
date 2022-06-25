import {
  Grid,
  Container,
  Avatar,
  Stack,
  Paper,
  Typography,
} from '@mui/material'
import { AxiosError } from 'axios'
import { GetServerSideProps } from 'next'
import api, { User } from '@/api'
import { requestOption } from '@/api/requestOption'
import CompetitionPageBase from '@/components/competitionPageBase'
import { MarkDown } from '@/components/markdown'

type DatePaperProps = {
  startAt: string
  endAt: string
}

type AuthorPaperProps = User

type RulePaperProps = {
  rule: string
}

type Props = {
  id: string
  name: string
  datePaperProps: DatePaperProps
  authorPaperProps: AuthorPaperProps
  rulePaperProps: RulePaperProps
}

// 開催期間
const DatePaper = (props: DatePaperProps) => {
  let startStr = new Date(props.startAt).toLocaleString()
  let endStr = new Date(props.startAt).toLocaleString()
  return (
    <Paper
      sx={{
        p: 2,
        display: 'flex',
        flexDirection: 'column',
        height: 120,
      }}
    >
      <Typography variant='h6' color='primary' gutterBottom>
        開催期間
      </Typography>
      <Stack direction='row' spacing={2} alignItems='center'>
        <Typography>
          {startStr} ~ {endStr}
        </Typography>
      </Stack>
    </Paper>
  )
}

// 著者情報
const AuthorPaper = (props: AuthorPaperProps) => {
  return (
    <Paper
      sx={{
        p: 2,
        display: 'flex',
        flexDirection: 'column',
        height: 120,
      }}
    >
      <Typography variant='h6' color='primary' gutterBottom>
        作問者
      </Typography>
      <Stack direction='row' spacing={2} alignItems='center'>
        <Avatar alt={props.name} src={props.iconUrl} />
        <span>{props.name}</span>
      </Stack>
    </Paper>
  )
}

// ルール説明
const RulePaper = (props: RulePaperProps) => {
  return (
    <Paper
      sx={{
        p: 2,
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      <Typography variant='h6' color='primary' gutterBottom>
        ルール説明
      </Typography>
      <MarkDown content={props.rule}></MarkDown>
    </Paper>
  )
}

const Competition = (props: Props) => {
  return (
    <CompetitionPageBase
      competitionId={props.id}
      competitionName={props.name}
      competitionRegisterd={true}
      competitionPageOption='overview'
    >
      <Container maxWidth='lg' sx={{ mt: 4, mb: 4 }}>
        <Grid container spacing={2}>
          {/* 開催期間 */}
          <Grid item xs={12} md={8}>
            <DatePaper {...props.datePaperProps}></DatePaper>
          </Grid>
          {/* 作問者 */}
          <Grid item xs={12} md={4}>
            <AuthorPaper {...props.authorPaperProps}></AuthorPaper>
          </Grid>
          {/* ルール */}
          <Grid item xs={12} md={12}>
            <RulePaper {...props.rulePaperProps}></RulePaper>
          </Grid>
        </Grid>
      </Container>
    </CompetitionPageBase>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const competitionId = ctx.params?.id as string
  let authorPaperProps: AuthorPaperProps = {
    name: '',
    iconUrl: '',
  }
  let datePaperProps: DatePaperProps = {
    startAt: '',
    endAt: '',
  }
  let propData: Props = {
    id: competitionId,
    name: '',
    authorPaperProps: authorPaperProps,
    rulePaperProps: {
      rule: '',
    },
    datePaperProps: datePaperProps,
  }
  try {
    const { data } = await api.getCompetition(competitionId, requestOption(ctx))
    propData.id = data.id
    propData.name = data.name
    // 開催期間
    propData.datePaperProps.startAt = data.startAt
    propData.datePaperProps.endAt = data.endAt
    // 作問者
    propData.authorPaperProps = data.author
    // ルール
    propData.rulePaperProps.rule = data.rule
  } catch (error) {
    const err = error as AxiosError
    console.log(err)
  }

  return {
    props: propData,
  }
}

export default Competition
