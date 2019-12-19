import React, { Component, Fragment } from 'react'
import gql from "graphql-tag";
import { Mutation } from 'react-apollo'
import QrReader from 'react-qr-scanner'
import Popup from "reactjs-popup";

/**
 * Query to create a shipment
 */
const FINISH_TRANSFER =
    gql`
    mutation ContinueDoubleSign($continueID: String!, $newHolderID: String!, $newHolderPassword: String!)
    {continueDoubleSign(request:
        {
            continueID: $continueID,
            newHolderID: $newHolderID,
            newHolderPassword: $newHolderPassword,
        }
    )}
`;

class FinishTransferShipmentPage extends Component {
    state = {
        continueID: '',
        newHolderID: '',
        newHolderPassword: '',
    }

    componentDidMount() {
        if (JSON.parse(localStorage.getItem('me'))) {
            const s = JSON.parse(localStorage.getItem('me'))["id"]
            this.setState({ newHolderID: s })
        }
    }

    handleScan(data) {
        // this.setState({
        //   result: data,
        // })
        console.log(data)
        this.setState({ continueID: data })
    }
    handleError(err) {
        console.error(err)
    }

    render() {
        if (!this.state.newHolderID) {
            return (
                <h2>Please login</h2>
            )
        }
        const { continueID, newHolderID, newHolderPassword } = this.state
        const previewStyle = {
            height: 240,
            width: 240,
        }
        return (
            <Mutation mutation={FINISH_TRANSFER} variables={{ continueID, newHolderID, newHolderPassword }}>
                {(continueDoubleSign, res) => {
                    const { data, loading, error, called } = res;
                    if (!called) {
                        return (
                            <section>
                                <h2>Finish Shipment Transfer </h2>
                                <form>
                                    <label>
                                        NewHolder ID
                                <input
                                            value={newHolderID}
                                            onChange={e => this.setState({ newHolderID: e.target.value })}
                                            type="text"
                                            required
                                        />
                                    </label>
                                    <br />
                                    <label>
                                        New Holder Password
                                <input
                                            value={newHolderPassword}
                                            onChange={e => this.setState({ newHolderPassword: e.target.value })}
                                            type="text"
                                            required
                                        />
                                    </label>
                                    <br />
                                    <label>
                                        Continue ID
                                <input
                                            value={continueID}
                                            onChange={e => this.setState({ continueID: e.target.value })}
                                            type="text"
                                            required
                                        />
                                    </label>
                                    <div>
                                        <Popup
                                            trigger={
                                                <button type="button" className="button">
                                                    {' '}
                                                    Open Modal{' '}
                                                </button>
                                            }
                                            modal
                                            closeOnDocumentClick>
                                            <div
                                                style={{ height: 500, width: 500, border: '1px solid #ccc' }}
                                            >
                                                <span> Scan continue id</span>
                                                <QrReader
                                                    delay={100}
                                                    style={previewStyle}
                                                    onError={this.handleError}
                                                    onScan={this.handleScan}
                                                />
                                            </div>
                                        </Popup>
                                        <p>{this.state.continueID}</p>
                                    </div>
                                </form>
                                <br />
                                <div>
                                    <button onClick={continueDoubleSign}>Finish Shipment Transfer</button>
                                </div>

                            </section>
                        )
                    }
                    if (loading) {
                        return <div>LOADING</div>;
                    }
                    if (error) {
                        return <div>ERROR <div>{data}</div></div>;
                    }

                    return (
                        <Fragment>
                            <h1>Transfer done</h1>
                        </Fragment>
                    )
                }}
            </Mutation>
        )
    }
}

export default FinishTransferShipmentPage
