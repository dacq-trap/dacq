import { Box, Typography, Grid, Avatar, Container } from '@mui/material'
import { User } from '@/api'

type Props = {
  memberList: Array<User>
}

const MemberList = (props: Props) => {
  return (
    <Box
      sx={{
        p: 2,
        mt: 4,
        mb: 4,
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: 'grey.50',
      }}
    >
      <Typography variant='h6' color='primary' gutterBottom>
        チームメンバーの一覧
      </Typography>
      <Container sx={{ mt: 4, mb: 4 }}>
        <Grid container spacing={2}>
          {props.memberList.map(({ name, iconUrl }) => (
            <Grid item key={name} xs={6} sm={3} md={2} sx={{ p: 1 }}>
              <Box display='flex' justifyContent='center'>
                <Avatar alt={name} src={iconUrl} />
              </Box>
              <Typography display='flex' justifyContent='center' noWrap>
                {name}
              </Typography>
            </Grid>
          ))}
        </Grid>
      </Container>
    </Box>
  )
}

export default MemberList
