import { useState } from 'react';
import { Card, Button } from 'react-bootstrap';
import "./HouseCard.css";

function HouseCard() {
    const [props,setProps] = useState(0)
    return (
        <div className="HouseCard">
            <Card>
                <Card.Img variant="top" src={props.img} />
                <Card.Body>
                    <Card.Title>{props.price}</Card.Title>
                    <Card.Text>
                        {props.comment}
                    </Card.Text>
                    <Card.Title>{props.address}</Card.Title>
                    <Button variant="primary">Read More</Button>
                </Card.Body>
            </Card>
        </div>
    );
}
export default HouseCard;  
