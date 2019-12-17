import React,{Component} from 'react'
import styled from '@emotion/styled'

class NotFound extends Component {
    render() {
        return (
            <div>
                <Message>404</Message>
                <img src="https://cdn.dribbble.com/users/135160/screenshots/6456597/go_lang_gopher_rahmen__2x.png"></img>
                <Message> There seems to be nothing here</Message>
            </div>
        )
    }

}

export default NotFound


const Message = styled.h1`
    text-align: center;
`


