import { Query } from 'react-apollo'
import gql from 'graphql-tag';
import React, { Component, Fragment } from 'react'

const TOQR =
    gql`
     query ToQR($dataID: String!)
     {toQR(data: $dataID)}
`;

class GetID extends Component {
    state = {
        dataID: ''
    }

    componentDidMount(){
        if (JSON.parse(localStorage.getItem('me'))) {
            const s = JSON.parse(localStorage.getItem('me'))["id"]
            this.setState({dataID: s})
        }
    }

    render() {
        if (!this.state.dataID) {
            return (
                <h2>Please login</h2>
            )
        }
        const { dataID } = this.state;
        return (
            <Query query={TOQR} variables={{ dataID }}>
                {({data, loading, error}) => {
                    if (!dataID) {
                        return <h1>Not logged in</h1>
                    }
                    if (loading) {
                        return <p>Loading</p>
                    }
                    if (error) {
                        return <p>Error</p>
                    }
                    const s = "data:image/png;base64," + data.toQR;
                    return (
                        <Fragment>
                            <h1>Your ID</h1>
                            <img src={s}/>
                        </Fragment>
                    )
                }}
            </Query>
        )
    }
}

export default GetID
