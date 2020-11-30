package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type StorageContainerResourceManagerId struct {
	ResourceGroup      string
	StorageAccountName string
	BlobServiceName    string
	ContainerName      string
}

func (id StorageContainerResourceManagerId) ID(subscriptionId string) string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/%s/containers/%s"
	return fmt.Sprintf(fmtString, subscriptionId, id.ResourceGroup, id.StorageAccountName, id.BlobServiceName, id.ContainerName)
}

func NewStorageContainerResourceManagerID(resourceGroup, accountName, blobServiceName, containerName string) StorageContainerResourceManagerId {
	return StorageContainerResourceManagerId{
		ContainerName:      containerName,
		StorageAccountName: accountName,
		BlobServiceName:    blobServiceName,
		ResourceGroup:      resourceGroup,
	}
}

func StorageContainerResourceManagerID(input string) (*StorageContainerResourceManagerId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	cache := StorageContainerResourceManagerId{
		ResourceGroup: id.ResourceGroup,
	}

	if cache.ContainerName, err = id.PopSegment("containers"); err != nil {
		return nil, err
	}

	if cache.BlobServiceName, err = id.PopSegment("blobServices"); err != nil {
		return nil, err
	}

	if cache.StorageAccountName, err = id.PopSegment("storageAccounts"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &cache, nil
}
