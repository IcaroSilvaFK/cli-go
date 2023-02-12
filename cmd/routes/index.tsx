
import {BrowserRouter,Routes as Router,Route} from 'react-router-dom'

export function Router(){
	return (
		<BrowserRouter>
			<Router>
			
<Route path='Dashboard' component={<Dashboard/>}/>

<Route path='Home' component={<Home/>}/>

<Route path='Main' component={<Main/>}/>
			</Router>
		</BrowserRouter>
	)
}
