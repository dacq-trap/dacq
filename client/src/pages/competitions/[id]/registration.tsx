import {
  Typography,
  Grid,
  Container,
  TextField,
  Box,
  Button,
} from '@mui/material'
import { AxiosError } from 'axios'
import { GetServerSideProps } from 'next'
import React, { useState } from 'react'
import api, { PostCompetitionTeamsRequest } from '@/api'
import { requestOption } from '@/api/requestOption'
import CompetitionPageBase from '@/components/competitionPageBase'

// 仮の構造
type Props = {
  competitionId: string
  competitionName: string
  myName: string
}

const Registration = (props: Props) => {
  const [teamName, setTeamName] = useState(props.myName)

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
    } catch (error) {
      const err = error as AxiosError
      console.log(err)
    }
  }

  return (
    <CompetitionPageBase
      competitionId={props.competitionId}
      competitionName={props.competitionName}
      competitionRegisterd={true}
      competitionPageOption='registration'
    >
      <Container maxWidth='lg' sx={{ mt: 4, mb: 4 }}>
        <Typography variant='h6' color='primary' gutterBottom>
          参加登録
        </Typography>
        <Grid container spacing={2}>
          {/* チーム名入力フォーム */}
          <Grid item xs={12} md={12}>
            <TextField
              id='teamName'
              name='teamNameForm'
              defaultValue={props.myName}
              label='チーム名'
              fullWidth
              variant='standard'
              onChange={(event) => setTeamName(event.target.value)}
              error={teamName.length < 1}
              helperText={teamName.length < 1 && '1文字以上入力してください'}
            />
          </Grid>
          {/* メンバーリスト（ユーザのみ） */}
          {/* 未実装 */}
          {/* 参加登録ボタン */}
          <Grid item xs={12} md={12}>
            <Box paddingY={2}>
              <form onClick={handleSubmit}>
                <Button
                  size='large'
                  variant='contained'
                  disabled={teamName.length < 1}
                  sx={{ maxWidth: 200, textAlign: 'center' }}
                >
                  参加登録する
                </Button>
              </form>
            </Box>
          </Grid>
        </Grid>
      </Container>
    </CompetitionPageBase>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const competitionId = ctx.params?.id as string

  let propData: Props = {
    competitionId: competitionId,
    competitionName: '',
    myName: '',
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
    propData.myName = data.name
  } catch (error) {
    const err = error as AxiosError
    console.log(err)
  }

  return {
    props: propData,
  }
}

export default Registration
