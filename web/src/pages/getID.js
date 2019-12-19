import { Query } from 'react-apollo'
import gql from 'graphql-tag';
import React, { Component, Fragment } from 'react'
import styled from '@emotion/styled'

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
                            <QRWrapper>
                                <h1>Your ID</h1>
                                <img src={s}/>
                                <h2>{dataID}</h2>
                            </QRWrapper>
                        </Fragment>
                    )
                }}
            </Query>
        )
    }
}

export default GetID


const QRWrapper = styled.div`
    border-style: solid;
    border-width: 5px;
    border-radius: 8px;
    text-align: center;
    max-width: 20%;
    margin: 50px;
    min-height: 500px;
    border-color: var(--main-color);
`
