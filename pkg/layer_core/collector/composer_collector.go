package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type ComposerCollector struct{}

func NewComposerCollector() *ComposerCollector {
	return &ComposerCollector{}
}

func (c *ComposerCollector) Satisfy(config map[string]interface{}, reference ast_contract.TokenReferenceInterface) (bool, error) {
	if !util.MapKeyExists(config, "composerPath") || !util.MapKeyIsString(config, "composerPath") {
		return false, layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("ComposerCollector needs the path to the composer.json file_supportive as string.")
	}
	if !util.MapKeyExists(config, "composerLockPath") || !util.MapKeyIsString(config, "composerLockPath") {
		return false, layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("ComposerCollector needs the path to the composer.lock file_supportive as string.")
	}
	if !util.MapKeyExists(config, "packages") || !util.MapKeyIsArrayOfStrings(config, "packages") {
		return false, layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("ComposerCollector needs the list of packages as strings.")
	}

	// TODO: implement go.mod parsing
	return false, nil

	//        try {
	//            $this->parser[$config['composerLockPath']] ??= new \Qossmic\Deptrac\Core\Layer\Collector\ComposerFilesParser($config['composerLockPath']);
	//            $parser = $this->parser[$config['composerLockPath']];
	//        } catch (RuntimeException $exception) {
	//            throw new CouldNotParseFileException('Could not parse composer files.', 0, $exception);
	//        }
	//        try {
	//            $namespaces = $parser->autoloadableNamespacesForRequirements($config['packages'], \true);
	//        } catch (RuntimeException $e) {
	//            throw InvalidCollectorDefinitionException::invalidCollectorConfiguration(\sprintf('ComposerCollector has a non-existent package defined. %s', $e->getMessage()));
	//        }
	//        $token = $reference->getToken()->toString();
	//        foreach ($namespaces as $namespace) {
	//            if (\str_starts_with($token, $namespace)) {
	//                return \true;
	//            }
	//        }
	//        return \false;
}
