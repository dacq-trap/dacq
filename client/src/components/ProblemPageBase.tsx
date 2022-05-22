import {
  Box,
  ButtonGroup,
  List,
  ListItemButton,
  ListItemText,
  Typography,
} from '@mui/material'
import Link from 'next/link'
import { useRouter } from 'next/router'
import Header from '@/components/header'

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
  const router = useRouter()
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
          <List style={{ display: 'flex', flexDirection: 'row' }}>
            {linkedButtonInfos.map(({ pageOption, text }) => (
              <ListItemButton
                key={pageOption}
                sx={{ maxWidth: 200, textAlign: 'center' }}
                selected={pageOption === props.problemPageOption}
                onClick={(event) => {
                  const href =
                    pageOption === 'overview'
                      ? `/problems/${props.problemId}`
                      : `/problems/${props.problemId}/${pageOption}`
                  router.push(href)
                }}
              >
                <ListItemText primary={text} />
              </ListItemButton>
            ))}
          </List>
        </Box>
      </Box>
      <Box paddingX={4}>{props.children}</Box>
    </>
  )
}

export default ProblemPageBase
