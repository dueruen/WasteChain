import React, { Fragment } from 'react';
import styled from '@emotion/styled'


export default function Navbar(props) {
    return (
        <Fragment>
            <nav>
                <NavWrapper>
                    <LeftSide>
                        <BrandWrapper href="/login">
                            <div >WasteChain</div>
                        </BrandWrapper>
                    </LeftSide>
                    <RightSide>
                        <LinkWrapper>
                            <NavLink href="#">Link</NavLink>
                        </LinkWrapper>
                        <LinkWrapper>
                            <NavLink href="/login">Login</NavLink>
                        </LinkWrapper>
                    </RightSide>
                </NavWrapper>
            </nav>
        </Fragment>
    )
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
