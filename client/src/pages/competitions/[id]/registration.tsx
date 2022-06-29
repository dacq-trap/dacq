import {
  Typography,
  Container,
  TextField,
  Box,
  Button,
  Paper,
  useThemeProps,
} from '@mui/material'
import { AxiosError } from 'axios'
import { GetServerSideProps } from 'next'
import React, { useState } from 'react'
import api, { PostCompetitionTeamsRequest, User } from '@/api'
import { requestOption } from '@/api/requestOption'
import CompetitionPageBase from '@/components/competitionPageBase'
import MemberList from '@/components/memberList'

type RegistrantProps = User

type Props = {
  competitionId: string
  competitionName: string
  isCompetitionRegistered: boolean
  registrantProps: RegistrantProps
}

const RegisteredComponent = () => {
  return (
    <Typography
      variant='h4'
      color='error'
      display='flex'
      justifyContent='center'
    >
      既に参加登録しています
    </Typography>
  )
}

const Registration = (props: Props) => {
  const [teamName, setTeamName] = useState(props.registrantProps.name)

  const handleSubmit = async (event: any) => {
    event.preventDefault()

    const postCompetitionTeamsRequest: PostCompetitionTeamsRequest = {
      name: teamName,
    }
    try {
      const { data } = await api.postCompetitionTeams(
        props.competitionId,
        postCompetitionTeamsRequest
      )
      console.log('hoge?')
    } catch (error) {
      const err = error as AxiosError
      console.log(err)
      console.log('huga?')
    }
  }

  return (
    <CompetitionPageBase
      competitionId={props.competitionId}
      competitionName={props.competitionName}
      isCompetitionRegistered={true}
      competitionPageOption='registration'
    >
      {/* 参加登録されてなければページを表示 */}
      {props.isCompetitionRegistered ? (
        <RegisteredComponent />
      ) : (
        <Paper
          sx={{
            p: 2,
            mt: 4,
            mb: 4,
          }}
        >
          <Typography
            variant='h6'
            color='primary'
            gutterBottom
            sx={{
              p: 1,
            }}
          >
            参加登録
          </Typography>
          <Container maxWidth='xl' sx={{ mt: 4, mb: 4 }}>
            <Box>
              {/* チーム名入力フォーム */}
              <TextField
                id='teamName'
                name='teamNameForm'
                defaultValue={props.registrantProps.name}
                label='チーム名'
                fullWidth
                variant='standard'
                onChange={(event) => setTeamName(event.target.value)}
                error={teamName.length < 1}
                helperText={teamName.length < 1 && '1文字以上入力してください'}
              />
              {/* メンバーリスト（ユーザのみ） */}
              <MemberList memberList={[props.registrantProps]} />
              {/* 参加登録ボタン */}
              <Box paddingY={2} display='flex' justifyContent='center'>
                <form onClick={handleSubmit}>
                  <Button
                    size='large'
                    variant='contained'
                    disabled={teamName.length < 1}
                  >
                    参加登録する
                  </Button>
                </form>
              </Box>
            </Box>
          </Container>
        </Paper>
      )}
    </CompetitionPageBase>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const competitionId = ctx.params?.id as string
  let propData: Props = {
    competitionId: competitionId,
    competitionName: '',
    isCompetitionRegistered: false,
    registrantProps: { name: '', iconUrl: '' },
  }

  // 既に参加登録しているか
  try {
    await api.getMyTeamInCompetition(competitionId, requestOption(ctx))
  } catch (error) {
    const err = error as AxiosError
    if (err.response && err.response.status === 404) {
      propData.isCompetitionRegistered = true
    } else {
      console.log(err)
    }
  }

  // getCompetition
  try {
    const { data } = await api.getCompetition(competitionId, requestOption(ctx))
    propData.competitionId = data.id
    propData.competitionName = data.name
  } catch (error) {
    const err = error as AxiosError
    console.log(err)
  }

  // getMe
  try {
    const { data } = await api.getMe(requestOption(ctx))
    propData.registrantProps = data
  } catch (error) {
    const err = error as AxiosError
    console.log(err)
  }

  return {
    props: propData,
  }
}

export default Registration
