import { Layout } from "antd";
import Nav from "./components/Nav";

export default function App() {
  const { Header, Footer, Content } = Layout;
  return (
    <Layout>
      <Header>
        <Nav />
      </Header>
      <Content>
        <div className="flex">
          123123123
        </div>
      </Content>
      <Footer>Footer</Footer>
    </Layout>
  );
}
