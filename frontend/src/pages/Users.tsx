// pages/Users.tsx
import { useEffect, useState } from 'react'
import { UserProps } from '../models/user'
import Layout from '../components/Layout'
import axios from 'axios'
import {
	Button,
	Table,
	TableBody,
	TableRow,
	TableHead,
	TableCell,
	TableFooter,
	TablePagination
} from '@material-ui/core'


const Users = () => {
	// 状態管理
	const [users, setUsers] = useState<UserProps[]>([])
	let ambassadorsUrl = 'ambassadors'

	// ページ情報のState
	const [page, setPage] = useState(0)
	const perPage = 10

	useEffect(() => {
		(
			async () => {
				const { data } = await axios.get(ambassadorsUrl)
				setUsers(data)
			}
		)()
	}, [])// 第二引数は第一引数に指定した関数の実行タイミングを決める
	      // 空を渡した場合、マウント・アンマウント時のみ第１引数の関数を実行

		return (
		<Layout>
			<Table className="table table-striped table-sm">
				<TableHead>...</TableHead>
				<TableBody>
					{users.slice(page * perPage, (page + 1) * perPage).map(user => {
						return (
							<TableRow key={user.id}>
								<TableCell>{user.id}</TableCell>
								<TableCell>{user.first_name} {user.last_name}</TableCell>
								<TableCell>{user.email}</TableCell>
								<TableCell>
									<Button
										variant="contained"
										color="primary"
										href={`users/${user.id}/links`}
										>View</Button>
								</TableCell>
							</TableRow>
						)
					})}
				</TableBody>
				<TableFooter>
					<TableRow>
						<TablePagination
							count={users.length}
							page={page}
							onPageChange={(e, newPage) => setPage(newPage)}
							rowsPerPageOptions={[]}
							rowsPerPage={perPage}
							></TablePagination>
					</TableRow>
				</TableFooter>
			</Table>
		</Layout>
	)
}

export default Users