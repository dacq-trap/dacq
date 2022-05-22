import { Box, ButtonGroup, Typography } from '@mui/material'
import Link from 'next/link'
import Header from './Header'
import LinkedButton from './LinkedButton'

// TODO: 追加する
type ProblemPageOption = 'overview' | 'leaderboard'

type Props = {
  problemId: string
  problemPageOption: ProblemPageOption
  children: React.ReactNode
}

type LinkedButtonInfo = {
  pageOption: ProblemPageOption
  text: string
}

const linkedButtonInfos: LinkedButtonInfo[] = [
  { pageOption: 'overview', text: '概要' },
  { pageOption: 'leaderboard', text: '順位表' },
]

const ProblemPageBase = (props: Props) => {
  return (
    <>
      <Header />
      <Box padding={2}>
        <Box paddingY={1}>
          <Link href={`/problems/${props.problemId}`} passHref>
            <Typography
              variant='h6'
              component='a'
              sx={{
                display: { xs: 'none', md: 'flex' },
                fontWeight: 500,
                color: 'inherit',
                textDecoration: 'none',
              }}
            >
              {props.problemId}
            </Typography>
          </Link>
        </Box>
        <Box paddingY={2}>
          <ButtonGroup
            variant='text'
            color='inherit'
            aria-label='inherit text button group'
          >
            {linkedButtonInfos.map(({ pageOption, text }) => (
              <LinkedButton
                href={
                  pageOption === 'overview'
                    ? `/problems/${props.problemId}`
                    : `/problems/${props.problemId}/${pageOption}`
                }
                text={text}
                key={pageOption}
                buttonProps={
                  pageOption === props.problemPageOption
                    ? { color: 'primary' }
                    : { color: 'inherit' }
                }
              />
            ))}
          </ButtonGroup>
        </Box>
      </Box>
      <Box paddingX={4}>{props.children}</Box>
    </>
  )
}

export default ProblemPageBase
