import { Card, Button } from 'react-bootstrap';

function HouseCard() {
    return (
        <div className="HouseCard">
            <Card>
                {/* <Card.Img variant="top" src={img1} /> */}
                <Card.Body>
                    <Card.Title>Card Title</Card.Title>
                    <Card.Text>
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras vitae molestie magna. Vivamus sed molestie enim, eu convallis mauris. Aliquam pharetra velit ac enim maximus, a commodo augue hendrerit. Phasellus at aliquam est
                    </Card.Text>
                    <Button variant="primary">Read More</Button>
                </Card.Body>
            </Card>
        </div>
    );
}
export default HouseCard;  