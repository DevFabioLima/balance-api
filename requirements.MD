# Requirements

- Deploy a Web Application that displays data from the Ethereum blockchain
    - Your Application should accept an Ethereum `address` and return the Addresses current Balance
        - Example Input URL
            - `trm.example.com/address/balance/0xc94770007dda54cF92009BFF0dE90c06F603a09f`
        - Example Output JSON
            
            ```json
            {
            	"balance": 0.0001365 
            }
            ```
            
    - Your Application should call [Infura.io](http://Infura.io) to obtain Ethereum Balance
        
        <aside>
        💡 API KEY: your-api-key
        
        </aside>
        
        - Details
            
            [Infura.io](https://infura.io/) provides an API interface for exploring and interacting with web3 technologies. In this exercise we will utilize Infura to retrieve data from the Ethereum blockchain!
            
            - Main Endpoint [`https://mainnet.infura.io/v3/your-api-key`](https://mainnet.infura.io/v3/your-api-key)
            - `eth_getBalance` - https://docs.infura.io/infura/networks/ethereum/json-rpc-methods/eth_getbalance
    - Your Application should be encapsulated in a Docker Image and stored on ECR or Docker Hub
    - Your Application should be hosted on AWS
- Operationalize Your Application using Terraform to run the Docker Image on a [Cloud Compute](https://hub.docker.com/) of your choice. Options:
    - AWS EKS
    - AWS ECS
    - AWS EC2
    - Others?
- Prepare for a live demonstration, which will include
    - Issuing a call to the API
    - Coding a new feature
    - Deploying the new feature through your CI/CD pipeline
- Share a link to your **GitHub repo** so we can review your code prior to the presentation
    
    

# **Follow-Up Questions**

- How can we make this deployment secure from bad actors?
- Is this deployment HA (highly available)? How can we improve and optimize for HA?
- How can we deploy changes to the Web Application 100+ times a day? What processes would you put in place?
- How would you scale out this web app to service thousands of customers per minute?
- How would you monitor this deployment in the event of an outage? What system would your alerts integrate with?

---

## Presentation Details

We appreciate your time on this presentation and we hope it offers more insight into our work at TRM! Here are some details to keep in mind:

**What to expect:** This is your chance to set the stage during a 1-hour presentation where you will be asked to present your design to the interview panel. Since this will be interactive, we recommend pausing throughout for interviewers to ask questions along the way. As the expert, your goal is to educate the panel on your design and demonstrate your knowledge and approach to designing large-scale systems. 

**Questions:** After reviewing this document, you’re encouraged to ask clarifying questions as needed to make your assumptions. To do this, please tag @Patrick Archambeau  in a comment in this document who will respond within 24hrs. (see: [how to comment in notion](https://www.notion.so/help/comments-mentions-and-reminders))

**Time Allocation:** There is no time limit for the preparation however, you're encouraged to set aside enough time to structure the problem, consider the requirements, document your design, and code.

**Materials:** Although not required, you’re welcome to use **basic visuals and diagrams** to showcase your ideas. For this or any written analysis or design, please feel free to use any tools at your disposal. Google Docs, Notion, or any other tool is OK. For all code written, you can package it via a **private** Github repo [preferred], or via a zip file with the relevant. We’re excited to see what you come up with! 

**Cloud Reimbursement:** Upon completing this presentation, please send the receipt and reimbursement details for the project’s cloud spend to talent@trmlabs.com and we will promptly reimburse you!

**Considerations:** The interview panel will dedicate the first 10 minutes of the presentation to reviewing your design. If you would like to reallocate this time to your presentation or Q&A at the end, please send your design document to the recruiting team within 24 hours of the presentation. 

**Evaluation criteria**

- Comprehension of design requirements and/or components
- Asks clarifying questions as needed
- Good judgment and reasoning on decision decisions, tradeoffs, and alternatives
- Adaptability to feedback
- Ability to execute a design from end-to-end
- Clear communication, written and verbal