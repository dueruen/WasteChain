import { Mutation } from 'react-apollo'
import gql from 'graphql-tag';
import React, { Component, Fragment } from 'react'
import ReactLoading from 'react-loading';
import { navigate } from "@reach/router"

/**
 * Mutation to login
 */
const LOGIN =
    gql`
    mutation Login($username: String!, $password: String!)
    {login(request:
        {
            username: $username,
            password: $password
        }
    ){token, id}}
`;

class Login extends Component {
    state = {
        username: '',
        password: ''
    }

    render() {
        const { username, password } = this.state
        return (
            <section>
                <h2>Login</h2>
                <Mutation mutation={LOGIN} variables={{ username, password }}>
                    {(login, res) => {
                        const { data, loading, error, called } = res;
                        if (!called) {
                            return (
                                <Fragment>
                                    <form>
                                        <label>
                                            Username
                                            <input
                                                value={username}
                                                onChange={e => this.setState({ username: e.target.value })}
                                                type="text"
                                                required
                                            />
                                        </label>
                                        <label>
                                            Password
                                            <input
                                                value={password}
                                                onChange={e => this.setState({ password: e.target.value })}
                                                type="password"
                                                required
                                            />
                                        </label>
                                    </form>
                                    <div>
                                        <button onClick={login}>Login</button>
                                    </div>
                                </Fragment>
                            )
                        }
                        if (loading) {
                            return <ReactLoading type={'spin'}color={'#8bb849'} height={'20%'} width={'20%'}/>;
                        }
                        if (error) {
                            return <div>ERROR</div>;
                        }

                        if (!data) {
                            return <div>Error cut not login</div>
                        } else {
                            const { login } = data;
                            const { token, id } = login;
                            if (!token) {
                                return <div>Error could not log in</div>
                            }
                            localStorage.setItem('me', JSON.stringify({ token: token, id: id }));

                            navigate('/')
                            window.location.reload(false) //Necessary to reload the navbar, to get the logged in features

                            return null

                        }
                    }}
                </Mutation>
            </section>
        )
    }
}

export default Login
