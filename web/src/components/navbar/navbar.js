import React, { Fragment, Component } from 'react';
import styled from '@emotion/styled'
import { navigate } from "@reach/router"
import './navbar.css'



export default class Navbar extends Component {

    renderRightSide = () => {
        if(localStorage.getItem('me')=== null) {
            return(
                <RightSide>
                    <LinkWrapper>
                        <NavLink href="/login">Login</NavLink>
                    </LinkWrapper>
                </RightSide>
            )
        }
        return(
            <RightSide>
                <LinkWrapper>
                    <div onClick={this.logOut}><NavLink href="/">Log Out</NavLink></div>
                </LinkWrapper>
                <LinkWrapper>
                    <NavLink href="/my-id/">My ID</NavLink>
                </LinkWrapper>
                <LinkWrapper>
                    <NavLink href="/my-shipments/">My Shipments</NavLink>
                </LinkWrapper>
                <LinkWrapper>
                    <NavLink href="/shipment/finishtransfer/">Finish Shipment Transfer</NavLink>
                </LinkWrapper>
                <LinkWrapper>
                    <NavLink href="/shipment/create/">Create Shipment</NavLink>
                </LinkWrapper>
            </RightSide>
        )
    }

    logOut = () => {
        localStorage.removeItem('me')
    }

    logoOnClick = () => {
        navigate('/')
    }

    render() {
    return (
        <Fragment>
            <nav>
                <NavWrapper>
                    <LeftSide>
                        <BrandWrapper>
                            <div className="logo" onClick={this.logoOnClick}>WasteChain</div>
                        </BrandWrapper>
                    </LeftSide>
                        {this.renderRightSide()}
                </NavWrapper>
            </nav>
        </Fragment>
    )
    }
}

const NavWrapper = styled.div`
    display: flex;
    justify-content: space-between;
    padding: 5px 10px 5px 20px;
    background-color: var(--main-color);
`

const LeftSide = styled.div`
    display: flex;
    align-items: center;
    justify-content: center;

`

const BrandWrapper = styled.div`
    font-size: 32px;
    color: #ffff;
    font-weight:bold;
`

const RightSide = styled.div`
    display: flex;
`

const LinkWrapper = styled.div`
    height: 22px;
    padding: 10px;
    border-bottom: 1px solid transparent;
    transition: border-bottom 0.5s;
    text-align: center;
`
const NavLink = styled.a`
    color: #ffff;
    text-decoration: none;
    transition: color 0.5s;
`
