import { GetServerSideProps } from 'next'
import ProblemPageBase from '@/components/problemPageBase'
import { DataGrid, GridColDef, GridValueGetterParams } from '@mui/x-data-grid';
const columns: GridColDef[] = [
  { field: 'id', headerName: 'ID', width: 70 },
  { field: 'rank', headerName: '順位', width: 70 },
  { field: 'firstName', headerName: 'First name', width: 130 },
  { field: 'lastName', headerName: 'Last name', width: 130 },
  {
    field: 'age',
    headerName: 'Age',
    type: 'number',
    width: 90,
  },
  {
    field: 'fullName',
    headerName: 'Full name',
    description: 'This column has a value getter and is not sortable.',
    sortable: false,
    width: 160,
    valueGetter: (params: GridValueGetterParams) =>
      `${params.row.firstName || ''} ${params.row.lastName || ''}`,
  },
];

const rows = [
  { id: 1, rank: 2, lastName: 'Snow', firstName: 'Jon', age: 35 },
  { id: 2, rank: 3, lastName: 'Lannister', firstName: 'Cersei', age: 42 },
  { id: 3, rank: 1, lastName: 'Lannister', firstName: 'Jaime', age: 45 },
  { id: 4, rank: 5, lastName: 'Stark', firstName: 'Arya', age: 16 },
  { id: 5, rank: 6, lastName: 'Targaryen', firstName: 'Daenerys', age: null },
  { id: 6, rank: 4, lastName: 'Melisandre', firstName: null, age: 150 },
  { id: 7, rank: 7, lastName: 'Clifford', firstName: 'Ferrara', age: 44 },
  { id: 8, rank: 9, lastName: 'Frances', firstName: 'Rossini', age: 36 },
  { id: 9, rank: 8, lastName: 'Roxie', firstName: 'Harvey', age: 65 },
];


// 仮の構造
type Props = {
  problemId: string
}

const Leaderboard = (props: Props) => {
  return (
    <ProblemPageBase
      problemId={props.problemId}
      problemPageOption='leaderboard'
    >
      順位表ページ
      <div style={{ height: 400, width: '100%' }}>
        <DataGrid
          rows={rows}
          columns={columns}
          pageSize={5}
          rowsPerPageOptions={[5]}
        />
      </div>
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

export default Leaderboard
