import React, { Component, Fragment } from 'react'
//import gql from "graphql-tag";
//import { Mutation } from 'react-apollo'
import QrReader from 'react-qr-scanner'
//import Popup from "reactjs-popup";

/**
 * Query to create a shipment
 */
// const TRANSFER_SHIPMENT =
//     gql`
//     mutation TransferShipment($shipmentID: String!, $ownerID: String!, $receiverID: String!, $location: String!, $password: String!)
//     {transferShipment(request:
//         {
//             shipmentID: $shipmentID,
//             ownerID: $ownerID,
//             receiverID: $receiverID,
//             location: $location,
//             password: $password,
//         }
//     ){error, QRCode, ContinueID}}
// `;


class StartTransferShipmentPage extends Component {
    state = {
        shipmentID: '', //this.props.shipmentID,
        ownerID: '',
        receiverID: '',
        location: '',
        password: '',
    }

    // componentDidMount() {
    //     if (JSON.parse(localStorage.getItem('me'))) {
    //         const s = JSON.parse(localStorage.getItem('me'))["id"]
    //         this.setState({ ownerID: s })
    //     }
    // }

    handleScan(data) {
        // this.setState({
        //   result: data,
        // })
        console.log(data)
        //this.setState({ receiverID: data })
    }
    handleError(err) {
        console.error(err)
    }

    render() {
        // if (!this.state.ownerID) {
        //     return (
        //         <h2>Please login</h2>
        //     )
        // }
        // const { shipmentID, ownerID, receiverID, location, password } = this.state
        // const previewStyle = {
        //     height: 240,
        //     width: 240,
        // }
        return (
            <div>
            <QrReader
              delay={300}
              onError={this.handleError}
              onScan={this.handleScan}
              //style={{ width: '200px' }}
            />
            {/* <p>{this.state.receiverID}</p> */}
          </div>
        )
        // return (
        //     <Mutation mutation={TRANSFER_SHIPMENT} variables={{ shipmentID, ownerID, receiverID, location, password }}>
        //         {(transferShipment, res) => {
        //             const { data, loading, error, called } = res;
        //             // if (!called) {
        //             //     return (
        //             //         <section>
        //             //             <div>
        //             //             <QrReader
        //             //                 delay={100}
        //             //                 style={{ width: '100%' }}
        //             //                 onError={this.handleError}
        //             //                 onScan={this.handleScan}
        //             //             />
        //             //             </div>
        //             //             <h2>Transfer Shipment</h2>
        //             //             <form>
        //             //                 <label>
        //             //                     Shipment ID
        //             //             <input
        //             //                         value={shipmentID}
        //             //                         onChange={e => this.setState({ shipmentID: e.target.value })}
        //             //                         type="text"
        //             //                         required
        //             //                     />
        //             //                 </label>
        //             //                 <br />
        //             //                 <label>
        //             //                     Owner ID
        //             //             <input
        //             //                         value={ownerID}
        //             //                         onChange={e => this.setState({ ownerID: e.target.value })}
        //             //                         type="text"
        //             //                         required
        //             //                     />
        //             //                 </label>
        //             //                 <br />
        //             //                 <label>
        //             //                     Receiver ID
        //             //             <input
        //             //                         value={receiverID}
        //             //                         onChange={e => this.setState({ receiverID: e.target.value })}
        //             //                         type="text"
        //             //                         required
        //             //                     />
        //             //                 </label>
        //             //                 <div>
        //             //                     {/* <Popup
        //             //                         trigger={
        //             //                             <button type="button" className="button">
        //             //                                 {' '}
        //             //                                 Open Modal{' '}
        //             //                             </button>
        //             //                         }
        //             //                         modal
        //             //                         closeOnDocumentClick>
        //             //                         <div
        //             //                             style={{ height: 500, width: 500, border: '1px solid #ccc' }}
        //             //                         >
        //             //                             <span> Scan receivers id</span>
        //             //                             <QrReader
        //             //                                 delay={100}
        //             //                                 style={previewStyle}
        //             //                                 onError={this.handleError}
        //             //                                 onScan={this.handleScan}
        //             //                             />
        //             //                         </div>
        //             //                     </Popup> */}
        //             //                     <p>{this.state.receiverID}</p>
        //             //                 </div>
        //             //                 <br />
        //             //                 <label>
        //             //                     Location
        //             //             <input
        //             //                         value={location}
        //             //                         onChange={e => this.setState({ location: e.target.value })}
        //             //                         type="text"
        //             //                         required
        //             //                     />
        //             //                 </label>
        //             //                 <br />
        //             //                 <label>
        //             //                     Password
        //             //             <input
        //             //                         value={password}
        //             //                         onChange={e => this.setState({ password: e.target.value })}
        //             //                         type="password"
        //             //                         required
        //             //                     />
        //             //                 </label>
        //             //             </form>
        //             //             <br />
        //             //             <div>
        //             //                 <button onClick={transferShipment}>Transfer Shipment</button>
        //             //             </div>

        //             //         </section>
        //             //     )
        //             // }
        //             if (loading) {
        //                 return <div>LOADING</div>;
        //             }
        //             if (error) {
        //                 return <div>ERROR</div>;
        //             }
        //             const s = "data:image/png;base64," + data.transferShipment.QRCode;
        //             return (
        //                 <Fragment>
        //                     <h1>Transfer started</h1>
        //                     <img src={s} />
        //                     <div>
        //                         <p>To finish the receiver has to scan and finish the transfer</p>
        //                     </div>
        //                 </Fragment>
        //             )
        //         }}
        //     </Mutation>
        // )
    }
}

export default StartTransferShipmentPage
